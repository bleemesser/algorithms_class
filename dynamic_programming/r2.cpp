#include <iostream>
#include <vector>
#include <unordered_map>
#include <chrono>

struct Prices
{
    int ore_robot;
    int clay_robot;
    std::pair<int, int> obsidian_robot;
    std::pair<int, int> geode_robot;
};

class State
{
public:
    int turn;
    int ore_robots;
    int clay_robots;
    int obsidian_robots;
    int geode_robots;
    int ore;
    int clay;
    int obsidian;
    int geode;

    State()
    {
        turn = 0;
        ore_robots = 1;
        clay_robots = 0;
        obsidian_robots = 0;
        geode_robots = 0;
        ore = 0;
        clay = 0;
        obsidian = 0;
        geode = 0;
    }

    void step()
    {
        ore += ore_robots;
        clay += clay_robots;
        obsidian += obsidian_robots;
        geode += geode_robots;
        turn++;
    }

    bool can_build_ore_robot(const Prices& p)
    {
        return ore >= p.ore_robot;
    }

    bool can_build_clay_robot(const Prices& p)
    {
        return ore >= p.clay_robot;
    }

    bool can_build_obsidian_robot(const Prices& p)
    {
        return ore >= p.obsidian_robot.first && clay >= p.obsidian_robot.second;
    }

    bool can_build_geode_robot(const Prices& p)
    {
        return ore >= p.geode_robot.first && obsidian >= p.geode_robot.second;
    }

    void build_ore_robot(const Prices& p)
    {
        ore -= p.ore_robot;
        ore_robots++;
        step();
    }

    void build_clay_robot(const Prices& p)
    {
        ore -= p.clay_robot;
        clay_robots++;
        step();
    }

    void build_obsidian_robot(const Prices& p)
    {
        ore -= p.obsidian_robot.first;
        clay -= p.obsidian_robot.second;
        obsidian_robots++;
        step();
    }

    void build_geode_robot(const Prices& p)
    {
        ore -= p.geode_robot.first;
        obsidian -= p.geode_robot.second;
        geode_robots++;
        step();
    }

    // create a key for the cache
    std::pair<int, int> key() {
        return std::make_pair(turn, ore_robots * 100000000 + clay_robots * 1000000 + obsidian_robots * 10000 + geode_robots * 100 + ore * 10 + clay * 10 + obsidian * 10 + geode);
    }
};

long cache_hits;
long states;
int solve(State s, Prices p, int turns, std::unordered_map<std::pair<int, int>, std::pair<int, int>>& cache)
{
    states++;
    if (s.turn == turns) // if we are at the end of the game, return the score
    {
        cache[s.key()] = std::make_pair(s.geode, s.turn);
        return s.geode;
    }

    if (cache.find(s.key()) != cache.end()) // if we have already solved this state, return the cached value. if we are at a later turn, return 0 because we can't do better than the cached value
    {
        cache_hits++;
        if (s.turn >= cache[s.key()].second)
        {
            return 0;
        }
    }

    int result = 0;

    if (s.can_build_geode_robot(p))
    {
        State s2 = s;
        s2.build_geode_robot(p);
        result = std::max(result, solve(s2, p, turns, cache));
    }
    if (s.can_build_obsidian_robot(p))
    {
        State s2 = s;
        s2.build_obsidian_robot(p);
        result = std::max(result, solve(s2, p, turns, cache));
    }
    if (s.can_build_clay_robot(p))
    {
        State s2 = s;
        s2.build_clay_robot(p);
        result = std::max(result, solve(s2, p, turns, cache));
    }
    if (s.can_build_ore_robot(p))
    {
        State s2 = s;
        s2.build_ore_robot(p);
        result = std::max(result, solve(s2, p, turns, cache));
    }

    State s2 = s;
    s2.step();
    result = std::max(result, solve(s2, p, turns, cache));

    cache[s.key()] = std::make_pair(result, s.turn);
    return result;
}

int main()
{
    int turns = 40;

    Prices p;
    p.ore_robot = 4;
    p.clay_robot = 2;
    p.obsidian_robot = std::make_pair(3, 7);
    p.geode_robot = std::make_pair(2, 7);

    State s = State();
    std::unordered_map<std::pair<int, int>, std::pair<int, int>> cache;
    auto start = std::chrono::high_resolution_clock::now();
    std::cout << solve(s, p, turns, cache) << std::endl;
    auto stop = std::chrono::high_resolution_clock::now();
    auto duration = std::chrono::duration_cast<std::chrono::milliseconds>(stop - start);
    std::cout << duration.count() << std::endl;
    std::cout << cache_hits << std::endl;
    std::cout << states << std::endl;
    return 0;
}
