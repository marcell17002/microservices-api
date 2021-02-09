<?php

namespace App\Http\Controllers;

use App\Category;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Validator;

class CategoryController extends Controller
{
    public function index()
    {
        $category = Category::all();
        return response()->json(['status' => 'success', 'data' => $category]);
    }

    public function show($id)
    {
        $category = Category::find($id);
        if (!$category) {
            return response()->json(['status' => 'error', 'message' => 'category not found'], 404);
        }

        return response()->json(['status' => 'success', 'data' => $category]);
    }

    public function create(Request $request)
    {
        $rules = [
            'name_category' => 'required|string'
        ];

        $data = $request->all();
        $validator = Validator::make($data, $rules);

        if ($validator->fails()) {
            return response()->json([
                'status' => 'error',
                'message' => $validator->errors()
            ], 400);
        }

        $category = Category::create($data);
        return response()->json(['status' => 'success', 'data' => $category]);
    }

    public function update(Request $request, $id)
    {
        $rules = [
            'name_category' => 'string'
        ];

        $data = $request->all();
        $validator = Validator::make($data, $rules);

        if ($validator->fails()) {
            return response()->json([
                'status' => 'error',
                'message' => $validator->errors()
            ], 400);
        }

        $category = Category::find($id);

        if (!$category) {
            return response()->json([
                'status' => 'error',
                'message' => 'category not found'
            ], 404);
        }

        $category->fill($data);
        $category->save();

        return response()->json(['status' => 'success', 'data' => $category]);
    }

    public function destroy($id)
    {
        $category = Category::find($id);

        if (!$category) {
            return response()->json(['status' => 'error', 'message' => 'category not found'], 404);
        }

        $category->delete();
        return response()->json(['status' => 'success', 'message' => 'category deleted']);
    }
}
