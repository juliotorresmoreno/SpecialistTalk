import React from 'react'
import ModalBS from 'react-bootstrap/Modal'

type ModalProps = {
  show: boolean
  title: React.ReactNode
  handleClose: () => void
} & React.PropsWithChildren

const Modal: React.FC<ModalProps> = ({
  show,
  title,
  children,
  handleClose,
}) => {
  return (
    <ModalBS show={show} onHide={handleClose}>
      <ModalBS.Header closeButton>
        <ModalBS.Title>{title}</ModalBS.Title>
      </ModalBS.Header>
      <ModalBS.Body>{children}</ModalBS.Body>
    </ModalBS>
  )
}

export default Modal
