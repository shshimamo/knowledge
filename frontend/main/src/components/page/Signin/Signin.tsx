import React, { useState } from 'react';
import Link from 'next/link'
import { useRouter } from 'next/router';
import { useAuthUsecase } from '@/usecase/user/usecase'
import styles from './Signin.module.css';

export const Signin = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const { signin } = useAuthUsecase();
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      await signin({ email, password });
      await router.push('/');
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <form className={styles.formContainer} onSubmit={handleSubmit}>
      <h1 className={styles.heading}>Login</h1>

      <div>
        <label htmlFor="email" className={styles.inputLabel}>Email</label>
        <input
          id="email"
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
          className={styles.inputField}
        />
      </div>

      <div>
        <label htmlFor="password" className={styles.inputLabel}>Password</label>
        <input
          id="password"
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
          className={styles.inputField}
        />
      </div>

      <div className={styles.buttonContainer}>
        <button type="submit" className={styles.submitButton}>
          Login
        </button>
        <Link href="/signup" className={styles.signInLink}>
          Sign Up
        </Link>
      </div>
    </form>
  );
}