import React, { useState } from 'react';
import Link from 'next/link'
import { useRouter } from 'next/router';
import { useAuthUsecase } from '@/usecase/auth/usecase'
import { useCreateUserMutation } from '@/api/main/mutation/__generated__/createUser.mutation.generated'

export const Signup = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [createUser, { loading, error }] = useCreateUserMutation();
  const { signup } = useAuthUsecase();
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      await signup({ email, password }, name, createUser);
      await router.push('/');
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <form className="max-w-sm mx-auto mt-20 space-y-8" onSubmit={handleSubmit}>
      <h1 className="text-2xl font-bold text-center">Sign Up</h1>

      <div>
        <label htmlFor="name" className="text-sm font-bold">Name</label>
        <input
          id="name"
          type="text"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
          className="w-full mt-2 p-2 border border-gray-300 rounded text-black"
        />
      </div>


      <div>
        <label htmlFor="email" className="text-sm font-bold">Email</label>
        <input
          id="email"
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
          className="w-full mt-2 p-2 border border-gray-300 rounded text-black"
        />
      </div>

      <div>
        <label htmlFor="password" className="text-sm font-bold">Password</label>
        <input
          id="password"
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
          className="w-full mt-2 p-2 border border-gray-300 rounded text-black"
        />
      </div>

      {error && <p className="text-red-500 mt-2">Error: {error.message}</p>}

      <div className="flex items-center justify-between">
        <button type="submit" disabled={loading} className="py-2 px-4 bg-blue-500 text-white font-bold rounded hover:bg-blue-600">
          {loading ? 'Loading...' : 'Sign Up'}
        </button>
        <Link href="/signin" className="text-blue-500 hover:text-blue-600">
          Sign In
        </Link>
      </div>
    </form>
  );
}