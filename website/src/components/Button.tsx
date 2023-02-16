import { Button as _Button } from 'react-bootstrap'
import styled from 'styled-components'

const Button = styled<typeof _Button>(_Button)`
  border-radius: 0;
  box-shadow: none;
  border: none;
  border-width: 0;
`

export default Button
