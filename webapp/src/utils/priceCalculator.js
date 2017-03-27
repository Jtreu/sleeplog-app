
const totalPrice = (price, currency) => {
  var currencySymbol = ''
  switch (currency) {
    case 'usd':
      currencySymbol = '$'
      break
    default:
      currencySymbol = '$'
  }

  if (currency === 'usd') {
    price = currencySymbol + '' + price
  }
  return price
}

export default {
  totalPrice
}
