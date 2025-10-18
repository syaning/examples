#include <iostream>
#include "velox/common/memory/Memory.h"
#include "velox/vector/BaseVector.h"
#include "velox/vector/ComplexVector.h"
#include "velox/vector/ConstantVector.h"
#include "velox/vector/FlatVector.h"

using namespace facebook::velox;

int main()
{
    // init memory pool
    memory::MemoryManager::initialize({});
    auto memoryManager = memory::MemoryManager::getInstance();
    auto rootPool = memoryManager->addRootPool("root");
    auto leafPool = rootPool->addLeafChild("leaf");

    // create id vector with values 1 to 7
    auto id = BaseVector::create(INTEGER(), 7, leafPool.get());
    auto *idValues = id->asFlatVector<int32_t>()->mutableRawValues();
    for (vector_size_t i = 0; i < id->size(); ++i)
    {
        idValues[i] = static_cast<int32_t>(i + 1);
    }

    // create constant vector with value 42
    auto con =
        BaseVector::createConstant(INTEGER(), Variant(42), 7, leafPool.get());

    // create day_of_week vector with string values
    auto dow = BaseVector::create(VARCHAR(), 7, leafPool.get());
    auto *dowValues = dow->asFlatVector<StringView>()->mutableRawValues();
    std::vector<std::string> days = {
        "monday",
        "tuesday",
        "wednesday",
        "thursday",
        "friday",
        "saturday",
        "sunday"};
    for (vector_size_t i = 0; i < dow->size(); ++i)
    {
        dowValues[i] = StringView(days[i]);
    }

    // create row vector with the above three vectors as children
    std::vector<std::string> names = {"id", "constant_42", "day_of_week"};
    std::vector<TypePtr> types = {INTEGER(), INTEGER(), VARCHAR()};
    auto rowType = ROW(std::move(names), std::move(types));
    std::vector<VectorPtr> children = {id, con, dow};
    auto rowVector = std::make_shared<RowVector>(
        leafPool.get(), rowType, nullptr, id->size(), std::move(children));

    std::cout << rowVector->toString() << std::endl;
    std::cout << rowVector->toString(0, rowVector->size()) << std::endl;

    return 0;
}