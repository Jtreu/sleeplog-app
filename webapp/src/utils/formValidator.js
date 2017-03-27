const passwordMinLength = 5

const passwordHasMinLength = (value) => {
  return value.length >= passwordMinLength
}

const passwordContainsLetter = (value) => {
  return value.toLowerCase().match(/[a-z]/i) || value.indexOf(' ') >= 0
}

const passwordContainsNumber = (value) => {
  return /\d/.test(value)
}

const passwordContainsUppercase = (value) => {
  return true
  // return /[A-Z]/.test(value)
}

const passwordContainsSpecial = (value) => {
  return true
  // return /[@~`!#$%^&*+=\-[\]\\';,/{}|\\":<>?]/g.test(value)
}

export {
  passwordMinLength,
  passwordHasMinLength,
  passwordContainsLetter,
  passwordContainsNumber,
  passwordContainsUppercase,
  passwordContainsSpecial
}
