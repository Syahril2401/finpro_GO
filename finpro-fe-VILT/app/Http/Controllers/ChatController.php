<?php

namespace App\Http\Controllers;

use App\Services\GoApiService;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Session;

class ChatController extends Controller
{
    public function __construct(protected GoApiService $api) {}

    /**
     * Handle a chat message from the user and return an AI reply.
     */
    public function send(Request $request): \Illuminate\Http\JsonResponse
    {
        $validated = $request->validate([
            'message' => ['required', 'string', 'max:1000'],
        ]);

        // Build learning profile string from the last stored result
        $lastResult     = Session::get('onboarding_result', []);
        $learningProfile = $this->buildProfileString($lastResult);

        try {
            $response = $this->api->chat($validated['message'], $learningProfile);

            return response()->json([
                'success' => true,
                'reply'   => $response['data']['reply'] ?? 'Maaf, saya tidak bisa menjawab saat ini.',
            ]);
        } catch (\Throwable $e) {
            return response()->json([
                'success' => false,
                'reply'   => 'Maaf, terjadi kesalahan. Coba lagi ya!',
            ], 500);
        }
    }

    private function buildProfileString(array $result): string
    {
        if (empty($result)) {
            return 'Profile not yet determined (survey not completed).';
        }

        $cat = function (int $score): string {
            if ($score <= 12) return 'Low';
            if ($score <= 18) return 'Medium';
            return 'High';
        };

        $planning       = $result['planning_score']        ?? 0;
        $timeManagement = $result['time_management_score'] ?? 0;
        $cognitive      = $result['cognitive_score']       ?? 0;
        $reflection     = $result['reflection_score']      ?? 0;
        $total          = $result['total_score']           ?? ($planning + $timeManagement + $cognitive + $reflection);

        return sprintf(
            "Planning: %s (%d/25) | Time Management: %s (%d/25) | Cognitive: %s (%d/25) | Reflection: %s (%d/25) | Total: %d/100",
            $cat($planning), $planning,
            $cat($timeManagement), $timeManagement,
            $cat($cognitive), $cognitive,
            $cat($reflection), $reflection,
            $total
        );
    }
}
