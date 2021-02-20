# ArcherDAO API

## Overview

### Strategies & Bankrolls

Today, we execute ETH -> ETH trades.
Single-asset accounting
No carry risk
Single-block, single-transaction execution

The Dispatcher smart contract manages these transactions using it’s bankroll.
Bankroll vs Flash Loans
Ownership of the bankroll

There may be cases where carry risk is acceptable (e.g. Maker Auctions).
Strategies that cover multiple blocks

What kinds of strategies and bankroll do we need to support moving ahead?
Supplier-provided bankroll
No significant bankroll required
Externally-provided bankroll, without carry risk
Externally-provided bankroll, with carry risk

## License

Apache-2.0
