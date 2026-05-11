<?php

namespace App\Http\Controllers;

use App\Services\GoApiService;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Session;
use Inertia\Inertia;
use Inertia\Response;

class OnboardingController extends Controller
{
    public function __construct(protected GoApiService $api) {}

    // ─── Step 1: Sanctuary ────────────────────────────────────────────────────

    public function sanctuary(): Response
    {
        return Inertia::render('Onboarding/Sanctuary');
    }

    // ─── Step 2: Questionnaire ────────────────────────────────────────────────

    public function questionnaire(): Response
    {
        // Fetch real questions from Go API (for question_ids)
        // If API fails, frontend has hardcoded questions as fallback
        try {
            $data = $this->api->getQuestions();
            $questions = $data['data'] ?? [];
        } catch (\Throwable) {
            $questions = [];
        }

        return Inertia::render('Onboarding/Questionnaire', [
            'questions' => $questions,
        ]);
    }

    // ─── Step 3: Submit answers → Go API ─────────────────────────────────────

    public function submit(Request $request): \Symfony\Component\HttpFoundation\Response
    {
        $validated = $request->validate([
            'answers'                => ['required', 'array', 'min:1'],
            'answers.*.question_id'  => ['required', 'string'],
            'answers.*.answer_value' => ['required', 'integer', 'min:1', 'max:5'],
        ]);

        try {
            $result = $this->api->submitAssessment($validated['answers']);

            // Cache result in session for the Result page
            Session::put('onboarding_result', $result['data'] ?? []);
            Session::put('survey_completed', true);

            return to_route('onboarding.result');
        } catch (\Throwable $e) {
            return back()->withErrors(['error' => 'Analysis failed: ' . $e->getMessage()]);
        }
    }

    // ─── Step 4: Result ───────────────────────────────────────────────────────

    public function result(): Response
    {
        // Pull result from session (set by submit())
        $result = Session::get('onboarding_result', []);

        return Inertia::render('Onboarding/Result', [
            'result' => $result,
        ]);
    }
}
