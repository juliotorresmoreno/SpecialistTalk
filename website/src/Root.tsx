/**
 * En este archivo se configuran todas las dependencias que se inician en la raíz del proyecto
 */

import React from 'react'
import App from './components/App'
import { BrowserRouter } from 'react-router-dom'
import { Provider } from 'react-redux'
import { store } from './store'
import { PersistGate } from 'redux-persist/lib/integration/react'
import { persistStore } from 'redux-persist'

let persistor = persistStore(store)

/**
 *
 */
const Root: React.FC<{}> = () => {
  return (
    <Provider store={store}>
      <PersistGate persistor={persistor}>
        <BrowserRouter>
          <App />
        </BrowserRouter>
      </PersistGate>
    </Provider>
  )
}

export default Root
