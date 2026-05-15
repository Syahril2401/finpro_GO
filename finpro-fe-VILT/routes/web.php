<?php

use App\Http\Controllers\AuthController;
use App\Http\Controllers\ChatController;
use App\Http\Controllers\DashboardController;
use App\Http\Controllers\OnboardingController;
use Illuminate\Support\Facades\Route;
use Inertia\Inertia;

// ─── Public Routes ─────────────────────────────────────────────────────────────
Route::get('/', fn () => Inertia::render('Landing'))->name('landing');

// ─── Lumora UI Pages (frontend demo — no auth required) ────────────────────────
Route::get('/lumora', fn () => Inertia::render('LumoraDashboard'))->name('lumora.dashboard');
Route::get('/lumora/planner', fn () => Inertia::render('LumoraPlanner'))->name('lumora.planner');
Route::get('/lumora/notes', fn () => Inertia::render('LumoraNotes'))->name('lumora.notes');
Route::get('/lumora/onboarding', fn () => Inertia::render('LumoraOnboarding'))->name('lumora.onboarding');

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

        // Dashboard Routes
        Route::get('/dashboard', [DashboardController::class, 'index'])->name('dashboard');
        Route::get('/dashboard/planner', [DashboardController::class, 'planner'])->name('planner');
        Route::get('/dashboard/notes', [DashboardController::class, 'notes'])->name('notes');
        Route::get('/dashboard/weekly-targets', [DashboardController::class, 'weeklyTargets'])->name('targets');
        Route::get('/dashboard/progress', [DashboardController::class, 'progress'])->name('progress');
        Route::get('/dashboard/settings', [DashboardController::class, 'settings'])->name('settings');
        
        // Recommendation Detail
        Route::get('/recommendations/{id}', [DashboardController::class, 'recommendationDetail'])->name('recommendations.detail');

        // AI Chatbot API endpoint
        Route::post('/chat', [ChatController::class, 'send'])->name('chat');
    });
});
