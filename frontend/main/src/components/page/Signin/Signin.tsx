import Link from 'next/link'
import { useRouter } from 'next/router'
import React, { useState } from 'react'

import { useAuthUsecase } from '@/usecase/user/usecase'

import styles from './Signin.module.css'

export const Signin = () => {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const authUsecase = useAuthUsecase()
  const router = useRouter()

  const signin = async () => {
    try {
      await authUsecase.signin({ email, password })
      await router.push('/knowledge_list')
    } catch (error) {
      console.error(error)
    }
  }

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    signin().catch((error) => {
      console.error(error)
    })
  }

  return (
    <form className={styles.formContainer} onSubmit={handleSubmit}>
      <h1 className={styles.heading}>Login</h1>

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
          Login
        </button>
        <Link href='/signup' className={styles.signInLink}>
          Sign Up
        </Link>
      </div>
    </form>
  )
}
