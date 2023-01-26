import { createRoot } from 'react-dom/client'
import Root from './Root'
import './styles/_00-root.scss'

const container = document.getElementById('app') as HTMLElement
const root = createRoot(container)

const app = Root({})

root.render(app)
