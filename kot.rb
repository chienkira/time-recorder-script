# frozen_string_literal: true

require 'net/http'
require 'json'
require 'optionparser'

KOT_API_ENDPOINT = 'https://s3.kingtime.jp/gateway/bprgateway'

def user
  uri = URI KOT_API_ENDPOINT
  res = Net::HTTP.post_form(uri,
                            page_id: 'account_verify',
                            account: @options[:user],
                            password: @options[:pass])
  raise StandardError, 'Login failed!' if bad_response? res

  res_data = JSON.parse(res.body)
  @user ||= {
    name: res_data['user_data']['user']['name'],
    user_token: res_data['user_data']['user']['user_token'],
    auth_token: res_data['user_data']['token']['token_b'],
    clock_in_id: res_data['user_data']['timerecorder']['record_button'][0]['id'],
    clock_out_id: res_data['user_data']['timerecorder']['record_button'][1]['id']
  }
end

def clock_in
  uri = URI KOT_API_ENDPOINT
  res = Net::HTTP.post_form(uri,
                            id: user[:clock_in_id],
                            user_token: user[:user_token],
                            token: user[:auth_token])
  raise StandardError, 'Clock in failed!' if bad_response? res

  puts "#{user[:name]} Clock in - 出勤 DONE"
end

def clock_out
  uri = URI KOT_API_ENDPOINT
  res = Net::HTTP.post_form(uri,
                            id: user[:clock_out_id],
                            user_token: user[:user_token],
                            token: user[:auth_token])
  raise StandardError, 'Clock out failed!' if bad_response? res

  puts "#{user[:name]} Clock out - 退勤 DONE"
end

def bad_response?(res)
  res_data = JSON.parse(res.body)
  res.code != '200' || res_data['result'] != 'OK'
end

# ============================================================

@options = {}
OptionParser.new do |opts|
  opts.banner = 'Usage: kot.rb [options]'

  opts.on('-u', '--user=USER_ID', ' Your login USER_ID') do |v|
    @options[:user] = v
  end

  opts.on('-p', '--pass=PASSWORD', 'Your login PASSWORD') do |v|
    @options[:pass] = v
  end

  opts.on('-a', '--action=ACTION', 'ACTION(IN or OUT)') do |v|
    @options[:action] = v
  end
end.parse!

case @options[:action].to_s.upcase!
when 'IN'
  clock_in
when 'OUT'
  clock_out
else
  puts 'Please choose IN or OUT action!
Run `ruby kot.rb -h` to show usage.'
end
