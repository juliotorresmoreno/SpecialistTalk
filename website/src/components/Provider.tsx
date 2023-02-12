import React from 'react'
import Websocket from './Websocket'

type ProviderProps = {} & React.PropsWithChildren

const Provider: React.FC<ProviderProps> = ({ children }) => {
  return <Websocket>{children}</Websocket>
}

export default Provider
