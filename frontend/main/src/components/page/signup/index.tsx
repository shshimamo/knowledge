import React, { useState } from 'react';
import { useAuthUsecase } from '@/usecase/auth/usecase'
import { useCreateUserMutation } from '@/api/main/mutation/__generated__/createUser.mutation'

export const Signup = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [createUser, { loading, error }] = useCreateUserMutation();
  const { signup } = useAuthUsecase(createUser);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      await signup({ email, password }, name);
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

      <button type="submit">Sign up</button>
    </form>
  );
}