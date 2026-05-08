<?php

namespace App\Http\Middleware;

use App\Services\GoApiService;
use Closure;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Session;
use Symfony\Component\HttpFoundation\Response;

/**
 * Ensures user has completed the onboarding survey before accessing the dashboard.
 * If they haven't, redirect them to the sanctuary page.
 */
class SurveyCompleted
{
    public function __construct(protected GoApiService $api) {}

    public function handle(Request $request, Closure $next): Response
    {
        // Cache the survey completion status in the session to avoid
        // an API call on every page load.
        if (!Session::has('survey_completed')) {
            $completed = $this->api->hasSurveyResult();
            Session::put('survey_completed', $completed);
        }

        if (!Session::get('survey_completed')) {
            return redirect()->route('onboarding.sanctuary')
                ->with('info', 'Selesaikan assessment terlebih dahulu untuk mengakses dashboard.');
        }

        return $next($request);
    }
}
