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
        return Inertia::render('Dashboard');
    }

    public function planner(): Response
    {
        return Inertia::render('Dashboard/Planner');
    }

    public function notes(): Response
    {
        return Inertia::render('Dashboard/Notes');
    }

    public function weeklyTargets(): Response
    {
        return Inertia::render('Dashboard/WeeklyTargets');
    }

    public function progress(): Response
    {
        return Inertia::render('Dashboard/Progress');
    }

    public function settings(): Response
    {
        return Inertia::render('Dashboard/Settings');
    }

    public function recommendationDetail($id): Response
    {
        return Inertia::render('RecommendationDetail', ['id' => $id]);
    }
}
