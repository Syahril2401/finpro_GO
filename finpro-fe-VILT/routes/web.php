<?php

use App\Http\Controllers\AuthController;
use App\Http\Controllers\ChatController;
use App\Http\Controllers\DashboardController;
use App\Http\Controllers\OnboardingController;
use Illuminate\Support\Facades\Route;
use Inertia\Inertia;

// ─── Public Routes ─────────────────────────────────────────────────────────────
Route::get('/', fn () => Inertia::render('Landing'))->name('landing');

// Auth
Route::get('/login',    [AuthController::class, 'loginForm'])->name('login');
Route::post('/login',   [AuthController::class, 'login']);
Route::get('/register', [AuthController::class, 'registerForm'])->name('register');
Route::post('/register',[AuthController::class, 'register']);
Route::post('/logout',  [AuthController::class, 'logout'])->name('logout');

// ─── Authenticated Routes (logged in, but may not have done survey) ───────────
Route::middleware('go.auth')->group(function () {

    // Onboarding (no survey check — users MUST be able to access these)
    Route::prefix('onboarding')->name('onboarding.')->group(function () {
        Route::get('/sanctuary',     [OnboardingController::class, 'sanctuary'])->name('sanctuary');
        Route::get('/questionnaire', [OnboardingController::class, 'questionnaire'])->name('questionnaire');
        Route::post('/submit',       [OnboardingController::class, 'submit'])->name('submit');
        Route::get('/result',        [OnboardingController::class, 'result'])->name('result');
    });

    // ─── Survey-gated Routes (must have completed survey) ──────────────────
    Route::middleware('survey.done')->group(function () {

        // Dashboard
        Route::get('/dashboard', [DashboardController::class, 'index'])->name('dashboard');

        // AI Chatbot API endpoint
        Route::post('/chat', [ChatController::class, 'send'])->name('chat');
    });
});
