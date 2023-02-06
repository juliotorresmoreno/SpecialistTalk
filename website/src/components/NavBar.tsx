import React, { useRef, useState } from 'react'
import Container from 'react-bootstrap/Container'
import Nav from 'react-bootstrap/Nav'
import Navbar from 'react-bootstrap/Navbar'
import { Link } from 'react-router-dom'
import useOnScreen from '../hooks/useOnScreen'
import Modal from './Modal'
import SingInForm from './SignInForm'
import SingUpForm from './SignUpForm'
import { ISession } from '../models/session'
import authSlice from '../features/auth'
import { useAppDispatch } from '../store/hooks'

type NavBarProps = {} & React.PropsWithChildren

const NavBar: React.FC<NavBarProps> = (props) => {
  const toggleBtnRef = useRef<HTMLButtonElement | null>()
  const isVisible = useOnScreen(toggleBtnRef)
  const [showModal, setShowModal] = useState<Link | null>(null)
  const dispatch = useAppDispatch()

  const toggleNav = () => {
    if (isVisible) toggleBtnRef.current?.click()
  }

  type Link = {
    title: React.ReactNode
    form: React.ReactNode
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
          <Nav className="me-auto">
            <Link to="/forum" className="nav-link" onClick={toggleNav}>
              Foro
            </Link>
          </Nav>
          <Nav>
            <Link
              to=""
              className="nav-link"
              onClick={() =>
                setShowModal({
                  title: (
                    <>
                      <span className="material-symbols-outlined">login</span>
                      &nbsp;Sing In
                    </>
                  ),
                  form: (
                    <SingInForm
                      onSuccess={(session: ISession) => {
                        setShowModal(null)
                        dispatch(authSlice.actions.setSession(session))
                      }}
                    />
                  ),
                })
              }
            >
              <span className="material-symbols-outlined">login</span>
              &nbsp;Sign In
            </Link>
          </Nav>
          <Nav>
            <Link
              to=""
              className="nav-link"
              onClick={() =>
                setShowModal({
                  title: (
                    <>
                      <span className="material-symbols-outlined">
                        app_registration
                      </span>
                      &nbsp;Sign Up
                    </>
                  ),
                  form: (
                    <SingUpForm
                      onSuccess={(session: ISession) => {
                        setShowModal(null)
                        dispatch(authSlice.actions.setSession(session))
                      }}
                    />
                  ),
                })
              }
            >
              <span className="material-symbols-outlined">
                app_registration
              </span>
              &nbsp;Sign up
            </Link>
          </Nav>
        </Navbar.Collapse>
        <Modal
          show={showModal !== null}
          title={showModal?.title ?? ''}
          handleClose={() => setShowModal(null)}
        >
          {showModal?.form}
        </Modal>
      </Container>
    </Navbar>
  )
}

export default NavBar
