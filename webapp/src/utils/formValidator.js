export const passwordMinLength = 5

export const passwordHasMinLength = (value) => {
  return value.length >= passwordMinLength
}

export const passwordContainsLetter = (value) => {
  return value.toLowerCase().match(/[a-z]/i) || value.indexOf(' ') >= 0
}

export const passwordContainsNumber = (value) => {
  return /\d/.test(value)
}

export const passwordContainsUppercase = (value) => {
  return true
  // return /[A-Z]/.test(value)
}

export const passwordContainsSpecial = (value) => {
  return true
  // return /[@~`!#$%^&*+=\-[\]\\';,/{}|\\":<>?]/g.test(value)
}
