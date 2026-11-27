# Antigravity Project Rules for CM-Beetle / Beetle UX Lab

## UI Typography, Styling & Terminology Rules (Beetle UX Lab)

For full component and styling guidelines, see `.agents/skills/ui-development/SKILL.md` and `ui/DESIGN_SYSTEM.md`.

- **Key-Value & Font Size**: Keys use `font-normal text-text-muted` (never bold); values use `font-extrabold text-text-main`. Minimum readable font size is `text-sm` (`1.1rem / 17.6px via ui/src/app/globals.css @theme`); avoid `text-[9px]`, `text-[10px]`, `text-[11px]`.
- **Text Casing & Terminology**: Use Title Case or Sentence Case only (no `uppercase` or `tracking-wider` on static labels; reserve `uppercase` for dynamic badges). Use `Nodes` / `Node Groups`, `2. Target Cloud Optimizer`, and `onpremiseInfraModel`.
- **Brand Palette & Light Mode Contrast**: Emerald (`#10b981`) / Teal (`#14b8a6`) only (forbid `cyan-*`, `purple-*`, `sky-*`). Use `--border-input` (`#cbd5e1`) and high-contrast text (`text-slate-800` or `text-text-main`) in light mode.
- **Design System Linting**: Always run `npm run lint` inside `ui/` after UI edits and resolve all `@shadcn/lint` violations.

## Go Development Philosophy & Coding Rules

1. **Core Principles**: `Simple > Complex` (`Clear > Clever`), `Design for Today`, Go SRP, `No Fallback` (validate explicitly and fail fast), `No Circular Dependencies`, and linear maintainability. See `.agents/rules/go-development.md` and `.github/instructions/go.instructions.md` for full standards.
2. **Code Block Comments**: Strictly 1 line maximum (1 extra line only for an explicit example or reference), written in concise English describing current behavior only (no history narratives).

## Pre-Implementation Scope and Trade-off Check

1. **Execute Directly for Clear Tasks**: For clear bug fixes, routine edits, or explicitly specified changes, proceed without extra ceremony.
2. **Clarify Before Coding on Structural Changes**: When a request adds a new API, alters `imdl`/`transx` schemas, or has multiple viable designs with ambiguous requirements, confirm core assumptions (at most 3 questions) and state trade-offs against `Simple > Complex` and `No Fallback` before editing code.
3. **Use `design-review` Skill**: Invoke `.agents/skills/design-review/SKILL.md` for formal pre-implementation architecture or PR design reviews.

## Documentation & Text Punctuation Rules

1. **Avoid Book/Typesetting Symbols (`—` and `·`)**:
   - Do **not** use em-dashes (`—`) or middle dots (`·`) when writing documentation, comments, UI text, or explanations, as they are not commonly typed by humans on a standard keyboard.
   - Replace `—` with standard colons (`:`), commas (`,`), parentheses (`(...)`), or standard hyphens (`-`).
   - Replace `·` with commas (`,`), slashes (`/`), or natural conjunctions (`및`, `와`, `과`, `and`).

## Pull Request & Auto-Merge Rules

1. **Manual Push, PR, and Merge by Default**:
   - For core logic, API, UI, and other general changes, `git push`, PR creation, `/approve`, and merging are handled manually by the user.
   - Do not push branches, open PRs, or post `/approve` unless the user explicitly asks.
2. **Wait for CI Checks Before `/approve` (Explicit Submodule Releases Only)**:
   - When the user explicitly requests automated PR, merge, and tagging (such as `imdl` or `transx` submodule releases), never post `/approve` immediately after creating a PR.
   - First monitor PR check runs until every required CI workflow finishes with `success`, and only then post `/approve` (`.github/workflows/auto-merge.yml` triggers on `issue_comment` and retries merge for only ~60 seconds).
3. **PR Title & Description**:
   - Use a conventional commit title (`<= 50` chars), since `.github/workflows/auto-merge.yml` sets `MERGE_COMMIT_MESSAGE: "pull-request-title"`.
   - Reuse the commit message body verbatim as the PR description unless asked otherwise.

