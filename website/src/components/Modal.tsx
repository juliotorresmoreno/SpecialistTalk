import React from 'react'
import { Button } from 'react-bootstrap'
import ModalBS from 'react-bootstrap/Modal'

type ModalProps = {
  show: boolean
  title: string
  handleClose: () => void
  handleSave: () => void
  buttons?: React.ReactNode
} & React.PropsWithChildren

const Modal: React.FC<ModalProps> = ({
  show,
  title,
  buttons,
  children,
  handleSave,
  handleClose,
}) => {
  return (
    <ModalBS show={show} onHide={handleClose}>
      <ModalBS.Header closeButton>
        <ModalBS.Title>{title}</ModalBS.Title>
      </ModalBS.Header>
      <ModalBS.Body>{children}</ModalBS.Body>
      {buttons ? <ModalBS.Footer>buttons</ModalBS.Footer> : null}
    </ModalBS>
  )
}

export default Modal
