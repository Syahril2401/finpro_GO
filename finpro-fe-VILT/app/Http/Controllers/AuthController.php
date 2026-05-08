<?php

namespace App\Http\Controllers;

use App\Services\GoApiService;
use Illuminate\Http\Request;
use Inertia\Inertia;
use Inertia\Response;

class AuthController extends Controller
{
    public function __construct(protected GoApiService $api) {}

    // ─── Show Login Form ───────────────────────────────────────────────────────

    public function loginForm(): Response
    {
        return Inertia::render('Auth/Login');
    }

    // ─── Process Login ─────────────────────────────────────────────────────────

    public function login(Request $request)
    {
        $validated = $request->validate([
            'email'    => ['required', 'email'],
            'password' => ['required', 'string'],
        ]);

        $result = $this->api->login($validated);

        if ($result['success'] ?? false) {
            return redirect()->route('dashboard');
        }

        return back()->withErrors([
            'email' => $result['message'] ?? 'Login gagal. Periksa email dan password kamu.',
        ])->withInput($request->only('email'));
    }

    // ─── Show Register Form ────────────────────────────────────────────────────

    public function registerForm(): Response
    {
        return Inertia::render('Auth/Register');
    }

    // ─── Process Register ──────────────────────────────────────────────────────

    public function register(Request $request)
    {
        $validated = $request->validate([
            'name'     => ['required', 'string', 'max:100'],
            'email'    => ['required', 'email'],
            'password' => ['required', 'string', 'min:8'],
        ]);

        $result = $this->api->register($validated);

        if ($result['success'] ?? false) {
            // Token saved in session by GoApiService::register via login
            // Auto-login after register
            $loginResult = $this->api->login([
                'email'    => $validated['email'],
                'password' => $validated['password'],
            ]);

            if ($loginResult['success'] ?? false) {
                return redirect()->route('onboarding.sanctuary');
            }
        }

        return back()->withErrors([
            'email' => $result['message'] ?? 'Registrasi gagal. Coba gunakan email lain.',
        ])->withInput($request->only('name', 'email'));
    }

    // ─── Logout ────────────────────────────────────────────────────────────────

    public function logout(): \Illuminate\Http\RedirectResponse
    {
        $this->api->logout();
        return redirect()->route('landing');
    }
}
