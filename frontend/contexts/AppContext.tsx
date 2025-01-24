import type React from "react";
import { createContext, useState, useContext, type ReactNode } from "react";
import type {
  User,
  Assignment,
  Grade,
  ApiResponse,
  UserResponse,
} from "../types";
import { useToast } from "@/hooks/use-toast";

const BACKEND_URL =
  process.env.NEXT_PUBLIC_BACKEND_URL || "http://localhost:8080";

interface AppContextType {
  currentUser: User | null;
  setCurrentUser: (user: User | null) => void;

  token: string;
  setToken: (token: string) => void;

  login: (email: string) => Promise<{ user: User; token: string }>;
  register: (
    email: string,
    name: string,
    role: string,
  ) => Promise<{ user: User; token: string } | null>;

  submitAssignment: (
    title: string,
    subject: string,
    content: string,
  ) => Promise<Assignment | null>;
  getAssignments: (filterBy?: string) => Promise<Assignment[]>;
  submitGrade: (
    assignment_id: string,
    grade: number,
    feedback: string,
  ) => Promise<Grade | null>;
  getGrades: (student_id: string) => Promise<Grade[]>;
}

const AppContext = createContext<AppContextType | undefined>(undefined);

export const AppProvider: React.FC<{ children: ReactNode }> = ({
  children,
}) => {
  const [currentUser, setCurrentUser] = useState<User | null>(null);
  const [token, setToken] = useState("");
  const { toast } = useToast();

  const handleError = (message: string) => {
    toast({
      variant: "destructive",
      title: "Error",
      description: message,
    });
  };

  const login = async (email: string): Promise<UserResponse> => {
    const response = await fetch(`${BACKEND_URL}/login?email=${email}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
    });

    const result: ApiResponse<UserResponse> = await response.json();

    if (response.status !== 200) {
      throw new Error(result.message || `Error: ${response.status}`);
    }

    setToken(result.data.token);
    setCurrentUser(result.data.user);
    return { user: result.data.user, token: result.data.token };
  };

  const register = async (
    email: string,
    name: string,
    role: string,
  ): Promise<UserResponse | null> => {
    const response = await fetch(`${BACKEND_URL}/users`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, name, role }),
    });

    const result: ApiResponse<UserResponse> = await response.json();

    if (response.status !== 201) {
      handleError(result.message || `Error: ${response.status}`);
      return null;
    }

    setToken(result.data.token);
    setCurrentUser(result.data.user);
    return { user: result.data.user, token: result.data.token };
  };

  const submitAssignment = async (
    title: string,
    subject: string,
    content: string,
  ): Promise<Assignment | null> => {
    const response = await fetch(`${BACKEND_URL}/assignment`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({ title, subject, content }),
    });

    const result: ApiResponse<Assignment> = await response.json();

    if (!response.ok) {
      toast({
        variant: "destructive",
        title: "Failed to submit assignment",
        description: result.message || `Error: ${response.status}`,
      });
      return null;
    }

    return result.data;
  };

  const getAssignments = async (filterBy?: string): Promise<Assignment[]> => {
    let url = `${BACKEND_URL}/assignment`;

    if (filterBy && filterBy !== "") {
      url = `${BACKEND_URL}/assignment?filterby=${filterBy}`;
    }

    const response = await fetch(url, {
      headers: { Authorization: `Bearer ${token}` },
    });
    const result: ApiResponse<Assignment[]> = await response.json();

    if (!response.ok) {
      toast({
        variant: "destructive",
        title: "Failed to fetch assignments",
        description: result.message || `Error: ${response.status}`,
      });
      return [];
    }

    return result.data;
  };

  const submitGrade = async (
    assignment_id: string,
    grade: number,
    feedback: string,
  ): Promise<Grade | null> => {
    const response = await fetch(`${BACKEND_URL}/grade`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({ assignment_id, grade, feedback }),
    });

    const result: ApiResponse<Grade> = await response.json();

    if (!response.ok) {
      toast({
        variant: "destructive",
        title: "Failed to submit grade",
        description: result.message || `Error: ${response.status}`,
      });
      return null;
    }

    return result.data;
  };

  const getGrades = async (student_id: string): Promise<Grade[]> => {
    const response = await fetch(`${BACKEND_URL}/grade/${student_id}`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    const result: ApiResponse<Grade[]> = await response.json();

    if (!response.ok) {
      toast({
        variant: "destructive",
        title: "Failed to fetch grades",
        description: result.message || `Error: ${response.status}`,
      });
      return [];
    }

    return result.data;
  };

  return (
    <AppContext.Provider
      value={{
        currentUser,
        setCurrentUser,
        token,
        setToken,
        login,
        register,
        submitAssignment,
        getAssignments,
        submitGrade,
        getGrades,
      }}
    >
      {children}
    </AppContext.Provider>
  );
};

export const useApp = () => {
  const context = useContext(AppContext);
  if (context === undefined) {
    throw new Error("useApp must be used within an AppProvider");
  }
  return context;
};
