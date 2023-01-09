# Model
model:
  rest_name: rendertemplate
  resource_name: rendertemplates
  entity_name: RenderTemplate
  package: ignis
  group: core/workflow
  description: |2-

    This API allows to render a provided generic go template and get back the
    results, after applying the provided values plus a bunch of additional
    information related to the product (like namespace, organizational metadata
    etc). The templating engine allows to call various helpers functions,
    and aims to be one to one compatible with Kubernetes templating, so you can
    use it to generate your Kubernetes template when they embeds additional
    security policies needing information like namespace, API URL etc, so they
    can be sent to an operator. (function `env`, `expandenv` and `getHostByName`
    are disabled).

    ##### String functions

    - `trim`: removes space from either side of a string.
    - `trimAll`: Remove given characters from the front or back of a string.
    - `trimSuffix`: Trim just the suffix from a string.
    - `trimPrefix`: Trim just the prefix from a string.
    - `upper`: Convert the entire string to uppercase.
    - `lower`: Convert the entire string to lowercase
    - `title`: Convert to title case.
    - `untitle`: Remove title casing. `untitle` "Hello World" produces hello world.
    - `repeat`: Repeat a string multiple times.
    - `substr`: Get a sub string from a string. It takes three parameters (start,
    end, string).
    - `nospace`: Remove all white space from a string.
    - `trunc`: Truncate a string (and add no suffix).
    - `abbrev`: Truncate a string with ellipses.
    - `abbrevboth`: Abbreviate both sides.
    - `initials`: Given multiple words, take the first letter of each word and
    combine.
    - `randAlphaNum`: Generate cryptographically secure random strings `(0-9azA-Z)`.
    - `randAlpha`: Generate cryptographically secure random strings `(a-zA-Z)`.
    - `randNumeric`: Generate cryptographically secure random strings `(0-9)`.
    - `randAscii`: Generate cryptographically secure random strings (printable ASCII
    chars).
    - `wrap`: Lets you specify the string to wrap with.
    - `contains`: Test to see if one string is contained inside of another.
    - `hasSuffix`: Test whether a string has a given suffix.
    - `hasPrefix`: Test whether a string has a given prefix.
    - `quote`: Wrap a string in double quotes.
    - `squote`: Wrap a string in single quotes.
    - `cat`: The cat function concatenates multiple strings together into one,
    separating them with spaces.
    - `indent`: Indents every line in a given string to the specified indent width.
    - `nindent`: Same as the `indent` function, but adds a new line to the
    beginning of the string.
    - `replace`: Perform simple string replacement.
    - `plural`: Pluralize a string.
    - `snakecase`: Convert string from `camelCase` to `snake_case`.
    - `camelcase`: Convert string from `snake_case` to `CamelCase`.
    - `kebabcase`: Convert string from `camelCase` to `kebab-case`.
    - `swapcase`: Swap the case of a string using a word based algorithm.
    - `shuffle`: Shuffle a string.
    - `regexMatch`: Returns true if the input string contains any match of the
    regular expression.
    - `mustRegexMatch`: Same as `regexMatch`, but returns an error to the template
    engine if there is a problem.
    - `regexFindAll`: Returns a slice of all matches of the regular expression in
    the input string.
    - `mustRegexFindAll`: Same as `regexFindAll`, but returns an error if there is a
    problem.
    - `regexFind`: Return the first (left most) match of the regular expression in
    the input string.
    - `mustRegexFindAll`: Same as `regexFind`, but returns an error if there is a
    problem.
    - `regexReplaceAll`: Returns a copy of the input string, replacing matches of
    the Regexp with the replacement string replacement.
    - `mustRegexReplaceAll`: Same as `regexReplaceAll`, but returns an error if
    there is a problem.
    - `regexReplaceAllLiteral`: Returns a copy of the input string, replacing
    matches of the Regexp with the replacement string replacement.
    - `mustRegexReplaceAllLiteral`: Same as `regexReplaceAllLiteral`, but returns an
    error if there is a problem.
    - `regexSplit`: Slices the input string into sub strings separated by the
    expression and returns a slice of the sub strings between those expression
    matches.
    - `mustRegexSplit`: Same as `regexSplit`, but returns an error if there is a
    problem.
    - `regexQuoteMeta`: Returns a string that escapes all regular expression
    meta characters inside the argument text.

    ##### String list functions
    - `join`: Join a list of strings into a single string, with the given separator.
    - `splitList`: Split a string into a list of strings (returns a list).
    - `split`: Split a string into a list of strings (returns a map).
    - `splitn`: Splits a string into a dictionary with a maximum length of N.
    - `sortAlpha`: Sort a list of strings into alphabetical (lexicographical) order.

    ##### Integer math functions
    - `add`: Sum numbers. Accepts two or more inputs.
    - `add1`: Increment by 1.
    - `sub`: Subtract two numbers.
    - `div`: Perform integer division.
    - `mod`: modulo.
    - `mul`: Multiply numbers. Accepts two or more inputs.
    - `max`: Return the largest of a series of integers.
    - `min`: Return the smallest of a series of integers
    - `floor`: Returns the greatest float value less than or equal to input value.
    - `round`: Returns a float value with the remainder rounded to the given number
    to digits after the decimal point.
    - `randInt`: Returns a random integer value from min (inclusive) to max
    (exclusive).

    ##### Integer slice functions
    - `until`: builds a range of integers.
    - `untilStep`: generates a list of counting integers. Allows you to define a
    start, stop, and step.
    - `seq`: Works like the bash seq command.

    ##### Float math functions
    - `addf`: Sum floats.
    - `add1f`: Increment by 1.
    - `subf`: Subtract floating numbers.
    - `divf`: Perform division.
    - `maxf`: Return the largest of a series of floats.
    - `minf`: Return the smallest of a series of floats.
    - `floor`: Returns the greatest float value less than or equal to input value.
    - `round`: Returns a float value with the remainder rounded to the given number
    to digits after the decimal point.

    ##### Date functions
    - `now`: The current date/time.
    - `ago`: Returns duration from `time.Now` in seconds resolution.
    - `date`: formats a date.
    - `dateInZone`: Same as date, but with a timezone.
    - `duration`: Formats a given amount of seconds as a `time.Duration.`
    - `durationRound`: Rounds a given duration to the most significant unit.
    - `unixEpoch`: Returns the seconds since the Unix epoch.
    - `dateModify`: Takes a modification and a date and returns the timestamp.
    - `mustDateModify`: Same as `dateModify`, but returns an error if there was a
    problem.
    - `htmlDate`: Formats a date for inserting into an HTML date picker input field.
    - `htmlDateInZone`: Same as `htmlDate`, but with a timezone.
    - `toDate`: converts a string to a date.
    - `mustToDate`: Same as `toDate`, but returns an error if there was a problem.

    ##### Default functions
    - `default`: Set a default value if second value is zero.
    - `empty`: returns true if the given value is considered empty, and false
    otherwise.
    - `coalesce`: takes a list of values and returns the first non-empty one.
    - `all`: takes a list of values and returns true if all values are non-empty.
    - `any`: takes a list of values and returns true if any value is non-empty.
    - `fromJson`: decodes a JSON document into a structure.
    - `mustFromJson`: Same as `fromJson`, but returns an error if there was a
    problem.
    - `toJson`: encodes an item into a JSON string.
    - `mustToJson`: Same as `toJson`, but returns an error if there was a problem.
    - `toPrettyJson`: encodes an item into a pretty (indented) JSON string.
    - `mustToPrettyJson`: Same as `toPrettyJson`, but returns an error if there was
    a problem.
    - `toRawJson`: Encodes an item into JSON string with HTML characters not
    escaped.
    - `mustToRawJson`: Same as `toRawJson`, but returns an error if there was a
    problem.
    - `ternary`: This is similar to the C ternary operator.

    ##### Encoding and decoding functions
    - `b64enc`: Encode to Base 64.
    - `b64dec`: decode Base 64.
    - `b32enc`: Encode to Base 32.
    - `b32dec`: decode Base 32.

    ##### Lists and list functions
    - `first`: Get the head item.
    - `mustFirst`: Same as `first`, but returns an error if there was a problem.
    - `rest`: get the tail of the list (everything but the first item).
    - `mustRest`: Same as `rest`, but returns an error if there was a problem.
    - `last`: Get the last item on a list.
    - `mustLast`: Same as `last`, but returns an error if there was a problem.
    - `initial`: Returns all but the last element.
    - `mustInitial`: Same as `initial`, but returns an error if there was a problem.
    - `append`: Append a new item to a list, creating a new list.
    - `mustAppend`: same as `append`, but returns an error if there was a problem.
    - `prepend`: Push an element onto the front of a list
    - `mustPrepend`: Same as `prepend`, but returns an error if there was a problem.
    - `reverse`: Produce a new list with the reversed elements of the given list.
    - `mustReverse`: Same as `reverse`, but returns an error if there was a problem.
    - `uniq`: Generate a list with all of the duplicates removed.
    - `mustUniq`: Same as `uniq` returns an error to the template engine if there is
    a problem.
    - `without`: filters items out of a list.
    - `mustWithout`: Same as `without`, but returns an error if there was a problem.
    - `has`: Test to see if a list has a particular element,
    - `mustHas`: Same as `has`, but returns an error if there was a problem.
    - `compact`: Accepts a list and removes entries with empty values.
    - `mustCompact`: Same as `compact`, but returns an error if there was a problem.
    - `slice`: get partial elements of a List
    - `mustSlice`: Same as `slice`, but returns an error if there was a problem.
    - `chunk`: split a list into chunks of given size.

    ##### Dictionaries and dictionary functions
    - `dict`: Creates dictionaries from a list of pairs.
    - `get`: Given a map and a key, get the value from the map.
    - `set`: add a new key/value pair to a dictionary.
    - `unset`: Given a map and a key, delete the key from the map.
    - `hasKey`: returns true if the given dictionary contains the given key.
    - `pluck`: Give one key and multiple maps, and get a list of all of the matches.
    - `dig`: traverses a nested set of dictionaries, selecting keys from a list of
    values.
    - `merge`: Merge two or more dictionaries into one, giving precedence to the
    destination dictionary.
    - `mustMerge`: Same as `merge`, but returns an error if there was a problem.
    - `mergeOverwrite`: Merge two or more dictionaries into one, giving precedence
    from right to left, effectively overwriting values in the destination
    dictionary.
    - `mustMergeOverwrite`: Same as `mergeOverwrite` but returns an error if there
    was a problem.
    - `keys`: Return a list of all of the keys in one or more dictionary types.
    - `pick`: The pick function selects just the given keys out of a dictionary,
    creating a new dictionary.
    - `omit`: similar to `pick`, except it returns a new dictionary with all the
    keys that
    do not match the given keys.
    - `values`: The values function is similar to keys, except it returns a new list
    with all the values of the source dictionary.
    - `deepCopy`: takes a value and makes a deep copy of the value.
    - `mustDeepCopy`: Same as `deepCopy`, but returns an error if there was a
    problem.

    ##### Type conversion functions
    - `atoi`: Convert a string to an integer.
    - `float64`: Convert to a `float64`.
    - `int`: Convert to an int at the system’s width.
    - `int64`: Convert to an `int64`.
    - `toDecimal`: Convert a Unix octal to a `int64`.
    - `toString`: Convert to a string.
    - `toStrings`: Convert a list, slice, or array to a list of strings.

    ##### Paths and file path functions
    - `base`: Return the last element of a path.
    - `dir`: Return the directory, stripping the last part of the path.
    - `clean`: Clean up a path.
    - `ext`: Return the file extension.
    - `isAbs`: check whether a path is absolute.
    - `osBase`: Return the last element of a file path, depending of OS (will be
    UNIX path).
    - `osClean`: Clean up a path, depending on the os (will be UNIX path).
    - `osExt`: Return the file extension, depending on the OS (will be UNIX path)
    - `osIsAbs`: check whether a file path is absolute, depending on the OS (will be
    UNIX path)

    ##### Flow control function
    - `fail`: Unconditionally returns an empty string and an error with the
    specified text.

    ##### UUID function
    - `uuidv4`: generates a UUID version 4.

    ##### OS function:
    - `env`: Not available. Will returns "disabled".
    - `expandenv`: not available. Will return "disabled".

    ##### Semantic version functions
    - `semver`: Parses a string into a Semantic Version.
    - `semverCompare`: returns true if the constraint matches, or false if it does
    not match.

    ##### Reflection function
    - `kindOf`: Return the kind of an object.
    - `kindIs`: Returns true if the given object is of the given kind.
    - `typeOf`: returns the underlying type of a value.
    - `typeIs`: Like `kindIs`, but for types.
    - `typeIsLike`: Works as `typeIs`, except that it also de-references pointers.
    - `deepEqual`: returns true if two values are deeply equal.

    ##### Cryptographic functions
    - `sha1sum`: Computes the SHA 1 digest of the given string.
    - `sha256sum`: Computes the SHA 256 digest of the given string.
    - `alder32sum`: Computes the Alder 32 digest of the given string.
    - `bcrypt`: Generates a `bcrypt` has of the given string.
    - `htpasswd`: function takes a username and password and generates a `bcrypt`
    hash of the password.
    - `randBytes`: Accepts a count N and generates a cryptographically secure random
    sequence of N bytes.
    - `genPrivateKey`: generates a new private key encoded into a PEM block.
    - `buildCustomCert`: Generate a TLS certificate from given base 64 encoded
    certificate and key.
    - `genCA`: generates a new, self-signed x509 certificate authority using a
    2048 bit RSA private key.
    - `genCAWithKey`: generates a new, self-signed x509 certificate authority using
    a given private key.
    - `genSelfSignedCert`: generates a new, self-signed x509 certificate using a
    2048 bit RSA private key.
    - `genSelfSignedCertWithKey`: generates a new, self-signed x509 certificate
    using a given private key.
    - `genSignedCert`: function generates a new, x509 certificate signed by the
    specified CA, using a 2048 bit RSA private key.
    - `genSignedCertWithKey`: generates a new, x509 certificate signed by the
    specified CA, using a given private key.
    - `encryptAES`: encrypts text with AES 256 CBC and returns a base 64 encoded
    string.
    - `decryptAES`: receives a base 64 string encoded by the AES 256 CBC algorithm
    and returns the decoded text.

    ##### Network function
    - `getHostByName`: Not available. Will return "disabled".
  aliases:
  - cook
  - rtpl

# Attributes
attributes:
  v1:
  - name: output
    description: Holds the rendered template.
    type: string
    exposed: true

  - name: parameters
    description: Contains the computed parameters.
    type: external
    exposed: true
    subtype: map[string]interface{}

  - name: template
    description: Template of the recipe.
    type: string
    exposed: true
