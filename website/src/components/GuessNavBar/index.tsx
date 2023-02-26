import React, { useRef } from 'react'
import Container from 'react-bootstrap/Container'
import Nav from 'react-bootstrap/Nav'
import Navbar from 'react-bootstrap/Navbar'
import { Link } from 'react-router-dom'
import { useLogout } from '../../services/auth'
import './style.css'

type NavBarProps = {} & React.PropsWithChildren

const GuessNavBar: React.FC<NavBarProps> = (props) => {
  const toggleBtnRef = useRef<HTMLButtonElement | null>()
  const { apply } = useLogout()

  const logout = () => apply()
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
          <Nav>
            <Link to="/profile" className="nav-link">
              <span className="material-symbols-outlined">account_circle</span>
              Profile
            </Link>
            <Link to="/settings" className="nav-link">
              <span className="material-symbols-outlined">settings</span>
              Settings
            </Link>
            <Link to="" className="nav-link" onClick={logout}>
              <span className="material-symbols-outlined">logout</span>
              Logout
            </Link>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  )
}

export default GuessNavBar
