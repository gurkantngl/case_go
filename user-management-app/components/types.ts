// types.ts
export interface User {
    id: number;
    username: string;
    email: string;
    // Add other user properties as needed
  }
  
  export interface ApiEndpoints {
    getAllUsers: string;
    getUserById: (id: number) => string;
    saveUser: string;
    updateUser: (id: number) => string;
    deleteUser: (id: number) => string;
  }
  