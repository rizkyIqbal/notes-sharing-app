export interface User {
    id: string,
    username: string;
    password?: string;
    createdAt: string;
    updatedAt: string;
}

export interface UserResponse {
    status: number;
    message: string;
    data?: User;
}