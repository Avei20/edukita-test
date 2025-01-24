"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useApp } from "../contexts/AppContext";
import type { Role } from "../types";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

interface RegistrationProps {
  onRegister: () => void;
}

export default function Registration({ onRegister }: RegistrationProps) {
  const [email, setEmail] = useState("");
  const [name, setName] = useState("");
  const [role, setRole] = useState<Role>("STUDENT");
  const [isLoading, setIsLoading] = useState(false);
  const { register } = useApp();

  const handleRegister = async () => {
    if (email && name && role) {
      setIsLoading(true);
      try {
        await register(email, name, role);
        onRegister();
      } catch (error) {
        console.error("Registration failed:", error);
        alert("Registration failed. Please try again.");
      } finally {
        setIsLoading(false);
      }
    } else {
      alert("Please fill in all fields");
    }
  };

  return (
    <div className="flex flex-col gap-4 max-w-sm mx-auto">
      <Input
        type="email"
        placeholder="Email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />
      <Input
        type="text"
        placeholder="Name"
        value={name}
        onChange={(e) => setName(e.target.value)}
      />
      <Select value={role} onValueChange={(value: Role) => setRole(value)}>
        <SelectTrigger>
          <SelectValue placeholder="Select role" />
        </SelectTrigger>
        <SelectContent>
          <SelectItem value="STUDENT">Student</SelectItem>
          <SelectItem value="TEACHER">Teacher</SelectItem>
        </SelectContent>
      </Select>
      <Button onClick={handleRegister} disabled={isLoading} className="w-full">
        {isLoading ? "Registering..." : "Register"}
      </Button>
    </div>
  );
}
