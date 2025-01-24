"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { AppProvider } from "../contexts/AppContext";
import TeacherDashboard from "@/components/TeacherDashboard";
import StudentDashboard from "@/components/StudentDashboard";
import Login from "@/components/Login";
import Registration from "@/components/Registration";

export default function Home() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [showRegistration, setShowRegistration] = useState(false);

  return (
    <AppProvider>
      <div className="container mx-auto p-4">
        <h1 className="text-2xl font-bold mb-4">
          Assignment Management System
        </h1>
        {!isLoggedIn ? (
          showRegistration ? (
            <>
              <Registration onRegister={() => setIsLoggedIn(true)} />
              <div className="flex justify-center mt-4">
                <Button onClick={() => setShowRegistration(false)}>
                  Back to Login
                </Button>
              </div>
            </>
          ) : (
            <>
              <Login onLogin={(val: boolean) => setIsLoggedIn(val)} />
              <div className="flex justify-center mt-4">
                <Button onClick={() => setShowRegistration(true)}>
                  Register New User
                </Button>
              </div>
            </>
          )
        ) : (
          <>
            <Button onClick={() => setIsLoggedIn(false)} className="mb-4">
              Logout
            </Button>
            <TeacherDashboard />
            <StudentDashboard />
          </>
        )}
      </div>
    </AppProvider>
  );
}
