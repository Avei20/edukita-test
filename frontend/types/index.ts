export type Role = "TEACHER" | "STUDENT";

export type Subject = "English" | "Math";

export interface User {
  id: string;
  email: string;
  name: string;
  role: Role;
}

export interface Assignment {
  id: string;
  title: string;
  subject: Subject;
  content: string;
  student_id: string;
}

export interface Grade {
  id: string;
  assignment_id: string;
  grade: number;
  feedback: string;
}

export interface UserResponse {
  token: string;
  user: User;
}

export interface ApiResponse<T> {
  success: boolean;
  status_code: number;
  message: string;
  data: T;
}
