---
name: design-review
description: "Pre-implementation architecture and PR design review against CM-Beetle Go, imdl, transx, and API standards (clarifying questions, red-team audit, and option trade-off matrix)."
---

# Design Review for CM-Beetle

Review non-trivial feature proposals, architectural changes, or PR diffs before implementation against CM-Beetle engineering rules in `.agents/rules/go-development.md`, `.github/instructions/go.instructions.md`, and `.github/copilot-instructions.md`.

## When to Use

- Adding new REST APIs, recommendation/migration flows, or background workers
- Modifying shared schemas or package boundaries (`imdl/`, `pkg/modelconv/`, `pkg/client/tumblebug/`, `transx/`, `analyzer/`)
- Choosing between 2 or more implementation approaches
- Skip this skill for trivial bug fixes, typo fixes, or fully specified mechanical edits

## Step 1: Clarify Critical Ambiguities (Only When Needed)

If core requirements, failure behavior, or backward-compatibility constraints are unclear:

1. Summarize the inferred goal and hidden assumptions in 2 to 3 lines.
2. Ask at most **3 decisive questions** that materially change the design (do not pad to 3 if fewer suffice).
3. Include a concrete default assumption for each question so the user can reply "proceed with defaults".

If requirements are already clear, skip directly to Step 2.

## Step 2: Red-Team Review Against CM-Beetle Standards

Evaluate the proposed design strictly against these criteria. List at most **5 real issues** ordered by severity (High, Medium, Low). Do not invent filler issues, and label any unverified premise as `[Assumption]`.

1. **Simplicity and Scope (`Simple > Complex`, `Design for Today`)**: Flag premature abstractions, speculative hooks, or wrapper layers not justified by current requirements.
2. **Deterministic Error Handling (`No Fallback`)**: Flag silent defaults, catch-all fallbacks, or swallowed errors that mask invalid inputs or upstream CB-Tumblebug failures.
3. **Package Boundaries (`SRP`, `No Circular Dependencies`)**: Check unidirectional flow across `pkg/api/rest/`, `pkg/core/`, `pkg/client/tumblebug/`, `pkg/modelconv/`, `imdl/`, `transx/`, and `analyzer/`.
4. **API, Model, and Submodule Impact**: Check `api/swagger.yaml` and `ui/swagger-beetle.yaml` sync impact, user-centric error messages (`model.SimpleErrorResponse`), and breaking changes (`!`) on `imdl` or `transx` modules.
5. **Failure and Recovery Behavior**: Check partial migration failures, retry safety/idempotency, and context cancellation/timeouts on external CSP or CB-Tumblebug calls.

### Output Table

| Severity | Finding | Violated Principle or Premise | Failure Scenario | Recommended Fix | Verification Method |
| :------: | ------- | ----------------------------- | ---------------- | --------------- | ------------------- |

## Step 3: Option Trade-Off Matrix (When Alternatives Exist)

When 2 or more viable designs exist, compare them without hedging and state a single recommendation.

| Criterion                             | Option A | Option B | Key Trade-Off |
| ------------------------------------- | -------- | -------- | ------------- |
| Simplicity and readability            | ...      | ...      | ...           |
| Backward compatibility (API / `imdl`) | ...      | ...      | ...           |
| Failure safety (`No Fallback`)        | ...      | ...      | ...           |
| Implementation and maintenance cost   | ...      | ...      | ...           |

Conclude with:

- **Recommended Option**: Single choice and why alternatives are rejected
- **Reversal Condition**: Specific condition under which the recommendation should flip
- **Next Implementation Checklist**: Concrete files and verification commands (`go build ./...`, `go test ./...`, `make lint`, `make swag`, `npm --prefix ui run lint`)
