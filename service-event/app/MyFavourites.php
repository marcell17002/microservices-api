<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class MyFavourites extends Model
{
    protected $table = 'my_favourites';

    protected $fillable = [
        'event_id', 'user_id'
    ];
    protected $casts = [
        'created_at' => 'datetime:Y-m-d H:m:s',
        'updated_at' => 'datetime:Y-m-d H:m:s'
    ];
    public function event()
    {
        return $this->belongsTo('App\Event');
    }
}
