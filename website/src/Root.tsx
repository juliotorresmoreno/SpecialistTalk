/**
 * En este archivo se configuran todas las dependencias que se inician en la raíz del proyecto
 */

import React from 'react'
import App from './components/App'
import { BrowserRouter } from 'react-router-dom'
import { Provider } from 'react-redux'
import store from './store'

/**
 *
 */
const Root: React.FC<{}> = () => {
  return (
    <Provider store={store}>
      <BrowserRouter>
        <App />
      </BrowserRouter>
    </Provider>
  )
}

export default Root
