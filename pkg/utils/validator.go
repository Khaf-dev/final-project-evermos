package utils

import "github.com/go-playground/validator/v10"

var Validator = validator.New() // Saya menggunakan Shared Instance agar kode lebih efisien dan rapih
