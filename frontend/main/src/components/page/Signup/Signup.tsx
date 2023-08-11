import React, { useState } from 'react';
import Link from 'next/link'
import { useRouter } from 'next/router';
import { useAuthUsecase } from '@/usecase/user/usecase'
import styles from './Signup.module.css';

export const Signup = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const { signup } = useAuthUsecase();
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      await signup({ email, password }, name);
      await router.push('/');
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <form className={styles.formContainer} onSubmit={handleSubmit}>
      <h1 className={styles.heading}>Sign Up</h1>

      <div>
        <label htmlFor="name" className={styles.inputLabel}>Name</label>
        <input
          id="name"
          type="text"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
          className={styles.inputField}
        />
      </div>


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
          Sign Up
        </button>
        <Link href="/signin" className={styles.signInLink}>
          Sign In
        </Link>
      </div>
    </form>
  );
}