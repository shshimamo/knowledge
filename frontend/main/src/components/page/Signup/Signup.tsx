import Link from 'next/link'
import { useRouter } from 'next/router'
import React, { useState } from 'react'

import { useAuthUsecase } from '@/usecase/user/usecase'

import styles from './Signup.module.css'

export const Signup = () => {
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const authUsecase = useAuthUsecase()
  const router = useRouter()

  const signup = async () => {
    try {
      await authUsecase.signup({ email, password }, name)
      await router.push('/')
    } catch (error) {
      console.error(error)
    }
  }
  const handleFormSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    signup().catch((error) => {
      console.error(error)
    })
  }

  return (
    <form className={styles.formContainer} onSubmit={handleFormSubmit}>
      <h1 className={styles.heading}>Sign Up</h1>

      <div>
        <label htmlFor='name' className={styles.inputLabel}>
          Name
        </label>
        <input
          id='name'
          type='text'
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
          className={styles.inputField}
        />
      </div>

      <div>
        <label htmlFor='email' className={styles.inputLabel}>
          Email
        </label>
        <input
          id='email'
          type='email'
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
          className={styles.inputField}
        />
      </div>

      <div>
        <label htmlFor='password' className={styles.inputLabel}>
          Password
        </label>
        <input
          id='password'
          type='password'
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
          className={styles.inputField}
        />
      </div>

      <div className={styles.buttonContainer}>
        <button type='submit' className={styles.submitButton}>
          Sign Up
        </button>
        <Link href='/signin' className={styles.signInLink}>
          Sign In
        </Link>
      </div>
    </form>
  )
}
