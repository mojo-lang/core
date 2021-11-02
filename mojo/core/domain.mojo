
/// For example, when naming a mail domain, the user should satisfy both the
/// rules of this memo and those in RFC-822.  When creating a new host name,
/// the old rules for HOSTS.TXT should be followed.  This avoids problems
/// when old software is converted to use domain names.
///
/// The following syntax will result in fewer problems with many
///
/// applications that use domain names (e.g., mail, TELNET).
///
/// ```
/// <domain> ::= <subdomain> | " "
/// <subdomain> ::= <label> | <subdomain> "." <label>
/// <label> ::= <letter> [ [ <ldh-str> ] <let-dig> ]
/// <ldh-str> ::= <let-dig-hyp> | <let-dig-hyp> <ldh-str>
/// <let-dig-hyp> ::= <let-dig> | "-"
/// <let-dig> ::= <letter> | <digit>
/// <letter> ::= any one of the 52 alphabetic characters A through Z in
/// upper case and a through z in lower case
/// <digit> ::= any one of the ten digits 0 through 9
/// ```
@format('labels:separate(".")') // join/split
type Domain {
    labels: [String] @1
}