const updateIconSource = (source) => {
  if (!source || source.indexOf('cloudinary') < 0) {
    return source
  }

  let indexOfUpload = source.indexOf('upload/')
  let pre = source.substring(0, indexOfUpload)
  let post = source.substring(indexOfUpload).substring(7)

  return `${pre}upload/c_scale,q_auto:eco,w_70/${post}`
}

const updateBannerSource = (source) => {
  if (!source || source.indexOf('cloudinary') < 0) {
    return source
  }

  let indexOfUpload = source.indexOf('upload/')
  let pre = source.substring(0, indexOfUpload)
  let post = source.substring(indexOfUpload).substring(7)

  return `${pre}upload/c_fill,h_170,q_auto:eco,w_1024/${post}`
}

export {
  updateIconSource,
  updateBannerSource
}
