import React, { useContext } from 'react'
import styled from 'styled-components'
import useToggle from '../../hooks/useToggle'
import Context from '../../contexts/SocialContext'
import Floating from './Floating'

const Container = styled.div`
  display: flex;
  flex: 1;
`

const ToggleButton = styled.div`
  cursor: pointer;
  flex: 1;
  text-align: center;
`

const FloatingContainer = styled.div`
  display: flex;
  flex: 1;
  flex-direction: column;
  position: absolute;
  height: calc(100vh - 106px);
  width: 100%;
  transform: translateY(-100%);
  background-color: var(--bs-gray-200);
`

type ContactsProps = {}

const Contacts: React.FC<ContactsProps> = () => {
  const [isOpen, toggle] = useToggle(true)
  const { setActiveChat } = useContext(Context)

  const onToggle = () => {
    if (isOpen) setActiveChat(null)
    toggle()
  }

  return (
    <Container>
      {isOpen ? (
        <FloatingContainer>
          <Floating />
        </FloatingContainer>
      ) : null}
      <ToggleButton onClick={onToggle}>Contactos</ToggleButton>
    </Container>
  )
}

export default Contacts
