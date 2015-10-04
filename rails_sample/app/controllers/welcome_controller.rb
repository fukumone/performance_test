class WelcomeController < ApplicationController

  def index
    render :json => "OK"
  end
end
