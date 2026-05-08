<?php

namespace App\Services;

use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Session;

class GoApiService
{
    protected string $baseUrl;

    public function __construct()
    {
        $this->baseUrl = config('services.go_api.url', env('GO_API_URL', 'http://localhost:8008'));
    }

    protected function withAuth(): \Illuminate\Http\Client\PendingRequest
    {
        $token = Session::get('go_token');
        return Http::withToken($token)->acceptJson();
    }

    protected function withoutAuth(): \Illuminate\Http\Client\PendingRequest
    {
        return Http::acceptJson();
    }

    // ─── Auth ────────────────────────────────────────────────────────────────

    public function register(array $data): array
    {
        $response = $this->withoutAuth()
            ->post("{$this->baseUrl}/api/auth/register", $data);

        return $response->json();
    }

    public function login(array $data): array
    {
        $response = $this->withoutAuth()
            ->post("{$this->baseUrl}/api/auth/login", $data);

        $json = $response->json();

        if (($json['success'] ?? false) && isset($json['data']['token'])) {
            Session::put('go_token', $json['data']['token']);
        }

        return $json;
    }

    public function logout(): void
    {
        Session::forget('go_token');
    }

    // ─── Assessment ───────────────────────────────────────────────────────────

    public function getQuestions(): array
    {
        return $this->withAuth()
            ->get("{$this->baseUrl}/api/assessment/questions")
            ->json();
    }

    public function submitAssessment(array $answers): array
    {
        return $this->withAuth()
            ->post("{$this->baseUrl}/api/assessment/submit", $answers)
            ->json();
    }

    /**
     * Check if the authenticated user has already completed the assessment.
     */
    public function hasSurveyResult(): bool
    {
        try {
            $resp = $this->withAuth()
                ->get("{$this->baseUrl}/api/assessment/status")
                ->json();
            return (bool) ($resp['data']['has_completed'] ?? false);
        } catch (\Throwable) {
            return false;
        }
    }

    /**
     * Fetch the last result summary for the authenticated user.
     */
    public function getLastResult(): array
    {
        try {
            $resp = $this->withAuth()
                ->get("{$this->baseUrl}/api/dashboard")
                ->json();
            return $resp['data']['last_assessment'] ?? [];
        } catch (\Throwable) {
            return [];
        }
    }

    /**
     * Send a chat message to the Gemini AI chatbot.
     */
    public function chat(string $message, string $learningProfile = ''): array
    {
        return $this->withAuth()
            ->post("{$this->baseUrl}/api/assessment/chat", [
                'message'          => $message,
                'learning_profile' => $learningProfile,
            ])
            ->json();
    }

    // ─── Dashboard ────────────────────────────────────────────────────────────

    public function getDashboard(): array
    {
        return $this->withAuth()
            ->get("{$this->baseUrl}/api/dashboard")
            ->json();
    }

    public function getPlanner(): array
    {
        return $this->withAuth()
            ->get("{$this->baseUrl}/api/dashboard/planner")
            ->json();
    }

    public function getNotes(): array
    {
        return $this->withAuth()
            ->get("{$this->baseUrl}/api/dashboard/notes")
            ->json();
    }

    public function getWeeklyTargets(): array
    {
        return $this->withAuth()
            ->get("{$this->baseUrl}/api/dashboard/weekly-targets")
            ->json();
    }

    public function getAiStrategies(): array
    {
        return $this->withAuth()
            ->get("{$this->baseUrl}/api/dashboard/ai-strategies")
            ->json();
    }

    public function getProgress(): array
    {
        return $this->withAuth()
            ->get("{$this->baseUrl}/api/dashboard/progress")
            ->json();
    }

    public function isLoggedIn(): bool
    {
        return Session::has('go_token');
    }
}
