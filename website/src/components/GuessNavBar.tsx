import React, { useRef, useState } from 'react'
import Container from 'react-bootstrap/Container'
import Nav from 'react-bootstrap/Nav'
import Navbar from 'react-bootstrap/Navbar'
import { Link } from 'react-router-dom'

type NavBarProps = {} & React.PropsWithChildren

const GuessNavBar: React.FC<NavBarProps> = (props) => {
  const toggleBtnRef = useRef<HTMLButtonElement | null>()

  return (
    <Navbar collapseOnSelect expand="lg" bg="purple" variant="dark">
      <Container>
        <Link className="navbar-brand" to="/">
          Home
        </Link>
        <Navbar.Toggle
          ref={toggleBtnRef as any}
          aria-controls="responsive-navbar-nav"
        />
        <Navbar.Collapse id="responsive-navbar-nav">
          <Nav className="me-auto"></Nav>
          <Nav></Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  )
}

export default GuessNavBar
