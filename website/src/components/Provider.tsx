import React from 'react'

type ProviderProps = {} & React.PropsWithChildren

const Provider: React.FC<ProviderProps> = ({ children }) => {
  return <>{children}</>
}

export default Provider
