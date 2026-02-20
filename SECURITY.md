# Security Policy

## License Compliance

This project and its dependencies are subject to license compliance requirements per organizational security policy.

### Current Status

⚠️ **License Policy Alert**: This project uses dependencies with BSD-2-Clause licenses, which require review and approval per organizational policy.

**Alert Reference**: SB-LICENSE-POLICY-VIOLATION::bsd-2-clause  
**Analyzer**: ANL-030-SBOMAnalyzer  
**Date Identified**: 2026-02-16

### Affected Dependencies

The following dependencies use BSD-2-Clause licenses:
- `github.com/gorilla/websocket v1.4.0` - WebSocket implementation
- `github.com/pkg/errors v0.8.1` - Error handling utilities
- `github.com/gopherjs/gopherjs` - Indirect dependency (testing)

### Compliance Requirements

This project must comply with the following standards:

1. **SOC 2 (CC6.1)**: Organizations must protect themselves from unauthorized information disclosure, including adherence to software licensing agreements.

2. **GDPR (Article 32)**: Requires a level of security appropriate to the risk. Violating software licenses can lead to security vulnerabilities, thereby compromising data protection.

3. **NIST SP 800-53 (CM-10)**: Software Usage Restrictions. The organization uses software and associated documentation in accordance with contract agreements and copyright laws.

4. **FedRAMP (CM-10)**: Software Usage Restrictions. Organizations must use software in accordance with licensing agreements and copyright laws.

5. **ISO 27001 (A.8.1.1)**: Inventory of Assets. Violation of license policy may indicate improper asset management and lack of control over organizational assets.

### Documentation

For complete license information, see:
- [THIRD-PARTY-LICENSES.txt](THIRD-PARTY-LICENSES.txt) - Complete list of all dependencies and their licenses
- [LICENSE](LICENSE) - This project's Apache-2.0 license

### Automated Compliance Checks

This repository includes automated license compliance checking via GitHub Actions:
- **Workflow**: `.github/workflows/license-compliance.yml`
- **Frequency**: On every PR, push to master, and weekly scheduled scans
- **Tool**: `go-licenses` by Google

### Remediation Options

If BSD-2-Clause licenses are not approved by organizational policy, consider:

1. **gorilla/websocket** alternatives:
   - `nhooyr.io/websocket` (MIT license)
   - `gobwas/ws` (MIT license)

2. **pkg/errors** alternatives:
   - Native Go 1.13+ error wrapping (standard library)
   - `github.com/cockroachdb/errors` (Apache-2.0)

3. **gopherjs** alternatives:
   - Update testing approach to avoid this indirect dependency

### Reporting Security Issues

If you discover a security vulnerability or license compliance issue, please report it to:
- **Email**: security@kaleidoscope-inc.com
- **Subject**: [SECURITY] Grammes - [Brief Description]

Please include:
- Description of the issue
- Steps to reproduce (if applicable)
- Potential impact
- Suggested remediation (if any)

### Security Updates

We are committed to addressing security and compliance issues promptly:
- **Critical issues**: Within 24 hours
- **High severity**: Within 7 days
- **Medium severity**: Within 30 days
- **Low severity**: Next release cycle

### Contact

For questions about security or license compliance:
- **Security Team**: security@kaleidoscope-inc.com
- **Legal Team**: legal@kaleidoscope-inc.com
- **Compliance Team**: compliance@kaleidoscope-inc.com

---

**Last Updated**: 2026-02-16  
**Policy Version**: 1.0
