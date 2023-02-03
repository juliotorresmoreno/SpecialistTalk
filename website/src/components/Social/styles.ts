import styled from 'styled-components'
import {
  Container as _ContainerBS,
  FormControl as _FormControlBS,
} from 'react-bootstrap'

export const Container = styled.div``

export const ContainerBS = styled(_ContainerBS)`
  display: flex;
  flex-direction: row-reverse;
  padding: 0;
`

export const Label = styled.div`
  padding: var(--spacing-v1);
  width: 100%;
  background-color: bisque;
  line-height: 1rem;
  text-align: center;
  cursor: pointer;
`
export const Floating = styled.div`
  width: 200px;
  position: absolute;
  height: calc(100vh - 200px);
  transform: translateY(-100%);
  background-color: white;
  border-color: #ccc;
  background-color: red;
`
export const FormControl = styled(_FormControlBS)`
  border: 0;
  border-radius: 0;
  box-shadow: none;
  line-height: calc(var(--spacing-v1) * 3);
  padding: 0 var(--spacing-v1);
  &:hover,
  &:focus {
    box-shadow: none;
    text-decoration: none;
  }
`

export const ContactContainer = styled.div`
  padding: var(--spacing-v1);
`

export const Contact = styled.div`
  height: 20px;
  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;
  cursor: pointer;
`

export const ChatContainer = styled.div`
  width: 100px;
  display: inline-block;

  line-height: calc(var(--spacing-v1) * 3);
  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;
  cursor: pointer;
`
