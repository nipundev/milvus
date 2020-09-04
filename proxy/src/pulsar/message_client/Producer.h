#pragma once

#include "pulsar/Producer.h"
#include "Client.h"
#include "src/grpc/gen-milvus/suvlim.pb.h"

namespace milvus {
namespace message_client {

using Producer = pulsar::Producer;
using ProducerConfiguration = pulsar::ProducerConfiguration;

class MsgProducer {
 public:
  MsgProducer(std::shared_ptr<pulsar::Client> &client, const std::string &topic,
              const ProducerConfiguration &conf = ProducerConfiguration());

  Result createProducer(const std::string &topic);
  Result send(const Message &msg);
  Result send(const std::string &msg);
  Result send(const milvus::grpc::InsertOrDeleteMsg &msg);
  Result send(const milvus::grpc::SearchMsg &msg);
  Result Send(const milvus::grpc::GetEntityIDsParam);
  Result close();

  const Producer &
  producer() const { return producer_; }

 private:
  Producer producer_;
  std::shared_ptr<pulsar::Client> client_;
  ProducerConfiguration config_;
};

}
}