import React, { useEffect, useRef, useState } from 'react'
import Container from 'react-bootstrap/Container'
import Nav from 'react-bootstrap/Nav'
import Navbar from 'react-bootstrap/Navbar'
import { Link, useLocation } from 'react-router-dom'
import classnames from 'classnames'
import useOnScreen from '../hooks/useOnScreen'

type NavBarProps = {} & React.PropsWithChildren

const NavBar: React.FC<NavBarProps> = (props) => {
  const toggleBtnRef = useRef<HTMLButtonElement | null>()
  const isVisible = useOnScreen(toggleBtnRef)
  const location = useLocation()
  const links0 = [
    ['/forum', 'Foro'],
    //['/recommends', 'Recomendado'],
    ['/social', 'Social'],
  ]
  const links1 = [
    ['/sign-in', 'Ingresar'],
    ['/sign-up', 'Registrarse'],
  ]
  const toggleNav = () => {
    if (isVisible) toggleBtnRef.current?.click()
  }
  const renderLinks = ([link, title]: string[]) => {
    const classNames = classnames({
      'nav-link': true,
      active: location.pathname === link,
    })
    return (
      <Link key={link} className={classNames} to={link} onClick={toggleNav}>
        {title}
      </Link>
    )
  }
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
          <Nav className="me-auto">{links0.map(renderLinks)}</Nav>
          <Nav>{links1.map(renderLinks)}</Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  )
}

export default NavBar
