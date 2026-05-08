<?php

namespace App\Http\Controllers;

use App\Services\GoApiService;
use Inertia\Inertia;
use Inertia\Response;

class DashboardController extends Controller
{
    public function __construct(protected GoApiService $api) {}

    public function index(): Response
    {
        $overview = $this->api->getDashboard();
        // Last assessment result — prefer session cache, fallback to API
        $result   = session('onboarding_result')
            ?: ($overview['data']['last_assessment'] ?? []);

        return Inertia::render('Dashboard', [
            'overview' => $overview['data'] ?? [],
            'result'   => $result,
        ]);
    }
}
