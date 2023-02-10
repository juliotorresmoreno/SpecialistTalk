import React, { useState } from 'react'
import WebsocketContext, { Handler } from '../../contexts/WebsocketContext'

const Provider: React.FC<React.PropsWithChildren> = ({ children }) => {
  const [handlers, setHandlers] = useState<Handler[]>([])

  const addHandler = (x: Handler) => {
    if (!handlers.includes(x)) {
      setHandlers([...handlers, x])
    }
  }

  const removeHandler = (x: Handler) => {
    if (handlers.includes(x)) {
      setHandlers(handlers.filter((handler) => handler !== x))
    }
  }

  return (
    <WebsocketContext.Provider value={{ handlers, addHandler, removeHandler }}>
      {children}
    </WebsocketContext.Provider>
  )
}

export default Provider
