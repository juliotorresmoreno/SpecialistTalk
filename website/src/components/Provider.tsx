import React from 'react'
import Websocket from './Websocket'
import SocialProvider from './Social/Provider'

type ProviderProps = {} & React.PropsWithChildren

const Provider: React.FC<ProviderProps> = ({ children }) => {
  return (
    <Websocket>
      <SocialProvider>{children}</SocialProvider>
    </Websocket>
  )
}

export default Provider
