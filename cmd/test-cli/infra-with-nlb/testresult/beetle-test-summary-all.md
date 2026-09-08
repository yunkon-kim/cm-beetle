# CM-Beetle integration test summary (with NLB)

> [!IMPORTANT]
> This document provides an overall summary of automated integration test results for all provider-region pairs.

## Execution details

- **Test Date**: September 8, 2026
- **Start Time**: 17:10:35 KST
- **End Time**: 17:17:44 KST
- **Total Execution Duration**: 7m46s
- **CM-Beetle Version**: v0.6.0+ (44cc606)
- **imdl Version**: v0.1.13+ (44cc606)
- **CB-Tumblebug Version**: v0.13.3
- **CB-Spider Version**: v0.13.4
- **CB-MapUI Version**: v0.13.7

## High-level test status

| Metric | Count | Description |
|--------|-------|-------------|
| **Total CSP Pairs** | **10** | Number of unique CSP-Region configurations evaluated |
| Passed CSP Pairs | 1 | Pairs where all test steps succeeded |
| Failed CSP Pairs | 0 | Pairs where at least one test step failed |
| Skipped CSP Pairs | 9 | Pairs that were disabled in config |
| **Total Test Steps** | **150** | Total individual endpoint tests triggered |
| Passed Steps | 15 | Individual tests that succeeded |
| Failed Steps | 0 | Individual tests that failed |

## Provider-specific summary

| Provider-Region | Status | Duration | Steps Passed | Details |
|-----------------|--------|----------|--------------|---------|
| **AWS-Seoul** | ✅ **PASS** | 7m40s | 15 / 15 | [View Report](beetle-test-results-aws.md) |

