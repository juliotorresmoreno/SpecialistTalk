import React, { useState } from 'react'
import useToggle from '../../hooks/useToggle'
import Floating from './Floating'
import { Label } from './styles'

const Contacts = () => {
  const [isOpen, toggle] = useToggle()

  return (
    <>
      {isOpen ? <Floating /> : null}
      <div className="contacts">
        <Label onClick={toggle}>chats</Label>
      </div>
    </>
  )
}

export default Contacts
