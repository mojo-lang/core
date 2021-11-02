///
/// An email address identifies an email box to which email messages are delivered.
///
/// The general format of an email address is local-part@domain,
/// and a specific example is jsmith@example.com. An address consists of two parts.
/// The part before the @ symbol (local-part) identifies the name of a mailbox.
/// This is often the username of the recipient, e.g., jsmith.
/// The part after the @ symbol (domain) is a domain name that represents the administrative realm for the mail box,
/// e.g., a company's domain name, example.com.
///
@format('{local_part}@{domain}')
type EmailAddress {
    ///
    local_part: String @1

    ///
    domain: Domain @2
}
