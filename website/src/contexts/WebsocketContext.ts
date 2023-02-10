import React from 'react'

export type Handler = {
  event: 'open' | 'message' | 'error' | 'close'
}

type DefaultSocialContext = {
  handlers: Handler[]
  addHandler(x: Handler): void
  removeHandler(x: Handler): void
}

const defaultSocialContext: DefaultSocialContext = {
  handlers: [],
  addHandler(x) {},
  removeHandler(x) {},
}

const SocialContext = React.createContext(defaultSocialContext)

export default SocialContext
