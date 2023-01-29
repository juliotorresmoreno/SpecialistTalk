import styled from 'styled-components'
import { Container as _ContainerBS } from 'react-bootstrap'

export const Container = styled.div`
  position: fixed;
  margin-top: 100vh;
  transform: translateY(-100%);
  width: 100%;
  # translateX(-100%);
  # margin-left: calc(100% - 50px);
`

export const ContainerBS = styled(_ContainerBS)`
  display: flex;
  flex-direction: row-reverse;
  padding: 0;
`
