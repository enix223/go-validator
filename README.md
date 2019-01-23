# go-validator: Validator for go web app

## Usage

```golang
// Define a validator
val := Validator{
    "username": []validators.ValidateHandler{
        // Blank validator for `username` field
        validators.NewBlankValidator(
            "blank", // validator identifier
            "username should not be blank", // error message
            username, // value to validate
            "",       // blank value
        ),
    },
    "password": []validators.ValidateHandler{
        // Blank validator
        validators.NewBlankValidator(
            "blank", gotext.Get("password should not be blank"), password, "",
        ),
        // Regex validator
        validators.NewRegExpValidator(
            "min_length", // validator identifier
            "password length should be greater or equal to 6", // error message
            password,                      // value to validate
            regexp.MustCompile("^.{6,}$"), // validate pattern
        ),
    },
}

res := val.Validate()
if !res.IsValid() {
    fmt.Println(res)
}
```