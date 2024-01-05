// UserDetail.tsx
import React from 'react';
import { User } from './types';

interface Props {
  user: User;
  onSave: (user: User) => void;
}

const UserDetail: React.FC<Props> = ({ user, onSave }) => {
  const handleSave = async () => {
    try {
      const response = await fetch('http://localhost:8000/users', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(user),
      });
      if (response.ok) {
        onSave(user);
      } else {
        console.error('Failed to save user');
      }
    } catch (error) {
      console.error('Error saving user:', error);
    }
  };

  return (
    <div>
      <h1>User Detail</h1>
      <form>
        <label>
          Username:
          <input type="text" value={user.username} />
        </label>
        <label>
          Email:
          <input type="email" value={user.email} />
        </label>
        <button onClick={handleSave}>{user.id ? 'Save' : 'Create'}</button>
      </form>
    </div>
  );
};

export default UserDetail;
